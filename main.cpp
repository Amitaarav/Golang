#include<bits/stdc++.h>
using namespace std;
vector<int> countSmaller(vector<int>nums) {
        int n = nums.size();
        vector<int> mpp;
        for(int i = 0;i<n;i++){
            mpp[i]=0;
        }
        for(int i = 0 ;i<n;i++){
            for(int j = i+1;j<n;j++){
                if(nums[j]<nums[i])
                {
                    mpp[i]++;
                }
            }
        }
        for(int i=0;i<n;i++){
            cout<<mpp[i]<<" ";
        }
        return mpp;
    }
int main()
{
    vector<int> nums = {5,2,6,1};
    countSmaller(nums);
    return 0;
}