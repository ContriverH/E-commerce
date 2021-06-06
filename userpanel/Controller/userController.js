const userDatabase = require("../Models/userModel");
const bcrypt = require("bcryptjs");
const jwt = require("jsonwebtoken");
const {userRegisterValidation}=require("../Config/userValidations.js");
/*
@desc To Create A User
@route POST /api/v1/user/registration
@model check userModels in Models Folder
*/

exports.createUser = async (req,res)=>{
    try{
        const {error} = await userRegisterValidation(req.body);
        if(error) return res.status(400).json({error:error.details[0].message});

        const salt = await bcrypt.genSalt(10);
        // Storing Hashed Password
        req.body.password = await bcrypt.hash(req.body.password,salt);

        await userDatabase.create(req.body);

        return res.status(201).json({
            success : true,
            message : "User Registration SuccessFull",
        });
    }catch (e) {
        return res.status(400).json({
            success : false,
            message : "Some Error Has Occurred While Creating New User Account"
        });
    }
}
/*
@desc For User Login
@route POST /api/v1/user/login
@model {
mobileNumber
password
}
*/
exports.userLogin = async (req,res)=>{
    const userFound = await userDatabase.findOne({mobileNumber : req.body.mobileNumber});
    if(!userFound) return res.status(400).json({
        success : false ,
        message : "Mobile Number Not Found"
    });
    //Validating Password
    const validatePassword = await bcrypt.compare(req.body.password , userFound.password);
    if(!validatePassword) return res.status(400).json({
        success : false ,
        message : "Wrong Password"
    });
    //Generating Token
    const token = await jwt.sign({
        id : userFound._id,
        role : userFound.role
    },process.env.SECRET_TOKEN);
    //Returning Token
    return res.status(200).json({
        success : true,
        jwt : token,
    });

}