const mongoose = require("mongoose");
/*
User Model (Subject To Change in Future) =
firstName
lastName
email
mobileNumber
address
password
dateOfJoining
role
*/

const userSchema = new mongoose.Schema({
    firstName : {
        type : String,
        required : true ,
        min : 3 ,
    },
    lastName : {
        type : String,
        required : true ,
        min : 2 ,
    },
    email : {
        type : String,
        unique : true
    },
    mobileNumber : {
        type : String,
        required : true,
        min : 10 ,
        max : 12 ,
        unique : true
    },
    address : {
        type : String,
    },
    password : {
        type : String,
        required : true ,
    },
    dateOfJoining : {
        type : Date ,
        default : Date.now
    },
    role : {
        type : String,
        default : "user"
    }
});

module.exports = mongoose.model("user",userSchema);