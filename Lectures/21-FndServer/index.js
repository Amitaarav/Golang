const express = require("express")
const app = express()
const port = 3000

app.use(express.json())
app.use(express.urlencoded({ extended: true }))
app.get("/", (req, res) => {
    res.status(200).send("Welcome to my server")
})

app.post("/post", (req, res) => {
    let myJson = req.body
    res.status(200).send(myJson)
})

app.post("/postform", (req, res) => {
    let myJson = req.body
    res.status(200).send(JSON.stringify(myJson))
})

app.listen(port, () => {
    console.log(`Server is running on port http://localhost:${port}`)
})