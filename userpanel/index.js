const express = require("express");
const app = express();
const router = require("./Route/routes");
const databaseConfig = require("./Config/databaseConfig");
const dotenv = require("dotenv");
dotenv.config({
    path : "./Config/configurations.env"
});

app.use(express.json())
app.use("/api/v1",router);

databaseConfig()
    .then(()=>console.log("Connected To Database"))
    .catch(()=>console.log("Failed To Connect To Database"));



const PORT = process.env.PORT || 8000;
app.listen(PORT , ()=> console.log(`Server Has Started On PORT ${PORT}`));