const express = require("express");
const router = express.Router();
const {createUser,userLogin} = require("../Controller/userController");

router.route("/user/registration")
    .post(createUser);
router.route("/user/login")
    .post(userLogin);

module.exports = router;