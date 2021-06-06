const joi = require("joi");


exports.userRegisterValidation = data =>{
    const schema = joi.object({
        firstName : joi.string().min(3).required(),
        lastName : joi.string().min(3).required(),
        email : joi.string().email(),
        mobileNumber : joi.string().min(10).max(12).required(),
        address : joi.string(),
        //Minimum 6 characters, at least one uppercase letter, one lowercase letter and one number
        password : joi.string().min(6).regex(RegExp("^(?=.*?[A-Z])(?=(.*[a-z]){1,})(?=(.*[\\d]){1,})(?=(.*[\\W]){1,})(?!.*\\s).{6,}")).required()
    });
    return schema.validate(data);
}