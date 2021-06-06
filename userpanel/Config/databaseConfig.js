const mongoose = require("mongoose");

const database = async ()=>{
    await mongoose.connect(process.env.DB,{
        useNewUrlParser : true ,
        useCreateIndex : true ,
        useUnifiedTopology : true,
        useFindAndModify : false
    });
}

module.exports = database;