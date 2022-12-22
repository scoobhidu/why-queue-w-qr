const { SchemaTypes } = require('mongoose');
const connection=require('../base/connection');
const Schema = connection.Schema;
const userSchema= new Schema({
    'TeacherID':{type:SchemaTypes.String,required:true,unique:true},
    'JWT':{type:SchemaTypes.Array}
});
const userModel= connection.model('QRaccess',userSchema);
module.exports = userModel;