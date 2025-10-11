import numpy as np
import matplotlib.pyplot as plt
from scipy.stats import binom
from scipy.optimize import minimize_scalar

# Given parameters
lot_size = 1000  # Lot size (N)
AQL = 0.02  # Acceptable Quality Level (2%)
LTPD = 0.05  # Lot Tolerance Proportional Limit (5%)
producer_risk = 0.02  # Producer's risk (alpha)
consumer_risk = 0.1  # Consumer's risk (beta)

# Function to calculate OC curve
def OC_curve(n, c, p):
    """Calculate OC curve value P_a(p) for a given sample size n, acceptance number c, and defect level p."""
    return sum(binom.pmf(k, n, p) for k in range(0, c + 1))

# Find optimal n and c to satisfy both producer's and consumer's risks
def find_sampling_plan(AQL, LTPD, producer_risk, consumer_risk):
    def objective(n_c_tuple):
        n, c = int(n_c_tuple[0]), int(n_c_tuple[1])
        Pa_AQL = OC_curve(n, c, AQL) - (1 - producer_risk)  # Producer's risk condition
        Pa_LTPD = OC_curve(n, c, LTPD) - consumer_risk  # Consumer's risk condition
        return Pa_AQL**2 + Pa_LTPD**2  # Minimize squared error to approximate satisfaction of both conditions

    # Constrain n and c to be within realistic ranges
    result = minimize_scalar(lambda x: objective((x, x / 2)), bounds=(1, 200), method="bounded")
    n = int(result.x)
    c = int(n / 2)
    return n, c

# Get sampling plan (n, c)
n, c = find_sampling_plan(AQL, LTPD, producer_risk, consumer_risk)

# Generate probabilities for OC curve over a range of defect levels
defect_levels = np.linspace(0, 0.1, 100)
Pa_values = [OC_curve(n, c, p) for p in defect_levels]

# Plot OC curve
plt.figure(figsize=(10, 6))
plt.plot(defect_levels * 100, Pa_values, label=f'OC Curve (n={n}, c={c})', color='b')
plt.axvline(AQL * 100, color='g', linestyle='--', label="AQL")
plt.axvline(LTPD * 100, color='r', linestyle='--', label="LTPD")
plt.xlabel("Defect Level (%)")
plt.ylabel("Probability of Acceptance (P_a)")
plt.title("Operating Characteristic (OC) Curve")
plt.legend()
plt.grid()
plt.show()

# AOQ curve
AOQ_values = [(Pa * p * (lot_size - n) / lot_size) for Pa, p in zip(Pa_values, defect_levels)]
AOQL = max(AOQ_values)

# Plot AOQ curve
plt.figure(figsize=(10, 6))
plt.plot(defect_levels * 100, AOQ_values, label=f'AOQ Curve (AOQL={AOQL:.4f})', color='purple')
plt.xlabel("Defect Level (%)")
plt.ylabel("Average Outgoing Quality (AOQ)")
plt.title("Average Outgoing Quality (AOQ) Curve")
plt.legend()
plt.grid()
plt.show()

# ATI curve
ATI_values = [n + (1 - Pa) * (lot_size - n) for Pa in Pa_values]

# Plot ATI curve
plt.figure(figsize=(10, 6))
plt.plot(defect_levels * 100, ATI_values, label="ATI Curve", color='orange')
plt.xlabel("Defect Level (%)")
plt.ylabel("Average Total Inspection (ATI)")
plt.title("Average Total Inspection (ATI) Curve")
plt.legend()
plt.grid()
plt.show()

n, c, AOQL
