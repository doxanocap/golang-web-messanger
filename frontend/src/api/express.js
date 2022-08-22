const express = require("express")
const app = express()

app.listen(8080)

app.use(express.json())

app.get("/12", (req,res) => {
  console.log(req.body);
})
