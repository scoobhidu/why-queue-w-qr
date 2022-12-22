const UserModel = require('../model/dbmodel');

// const {SUCCESS, SERVER_ERROR,RESOURCE_NOT_FOUND} = require('../../utils/constants/app_constants').STATUS_CODES;
module.exports = {
        // add(jwt,TeacherID){
        //     console.log(jwt,TeacherID);
        //     UserModel.findOne({"TeacherID":TeacherID},(err,doc)=>{
        //         if(err){
        //             console.log('DB error',err);
        //         }
        //         else{
        //             console.log('DB success',doc);
        //         }
        //     })

        //     }
        async add(jwt,TeacherID,date){
            console.log(TeacherID);
            await UserModel.findOne({'TeacherID':TeacherID}).exec()
            .then((doc)=>{
                if(!doc){
                    console.log("Not found")
                }
                else{
                    console.log("found",doc);
                    let obj={};
                    obj[jwt]=date;
                    doc.JWT.push(obj);
                    UserModel.findOneAndUpdate({'TeacherID':TeacherID},{JWT:doc.JWT}).exec()
                        .then((doc)=>{
                            if(!doc){
                                console.log("Not Found");
                            }
                            else{
                                console.log(doc);
                            }
            })
            .catch((err)=>{
                console.log(err);
            })
                }
            })
            .catch((err)=>{
                console.log(err);
                return err;
            })
        },
        async clear(userName){
            await UserModel.findOneAndUpdate({'TeacherID':userName},{JWT:[]}).exec()
            .then((doc)=>{
                if(!doc){
                    console.log("Not Found");
                }
                else{
                    console.log(doc);
                }
            })
            .catch((err)=>{
                console.log(err);
            })
        }
        }
    
    // async showUserWiseOrder(userid, orderid){
    //     console.log('Userid is ', userid);
    //     return await UserModel.findOne({"_id":userid}).populate('orderid').exec()
    //    },
    // async updateOrder(userid , orderid){
    //    return await UserModel.findByIdAndUpdate({_id:userid}, {orderid:orderid}).exec();
    //     // UserModel.findByIdAndUpdate({_id:userid}, {orderid:orderid}, (err)=>{
    //     //     if(err){

    //     //     }
    //     //     else{

    //     //     }
    //     // })
    // },

    // updatePWD(userObject, response){
    //     console.log('Inside Update ', userObject);
    //     UserModel.findOne({'email':userObject.email},(err,doc)=>{
    //         if(err){
    //             console.log('Error During Update in Login ', err);
    //             response.status(SERVER_ERROR).json({message:'Some Problem During Login '+err});
    //         }
    //         else if(doc && doc.email){
    //             let dbPassword = doc.password;
    //             let plainPassword = userObject.password;
    //             console.log('Correct Login ',userObject);
    //             if(encrypt.compareHash(plainPassword, dbPassword)){
    //                 // response.status(SUCCESS).json({message:message_reader.readMessage('user.welcome')+doc.name})
                    
    //                 userObject.newpwd = encrypt.hashPassword(userObject.newpwd);
    //                 console.log('Update New Password Hash ', userObject.newpwd);
    //                 UserModel.updateOne({'email':userObject.email},{password:userObject.newpwd},(err,doc)=>{
    //                     if(err){
    //                         console.log('Error in Updation ',err);
    //                         response.status(SERVER_ERROR).json({message:message_reader.readMessage('user.updatefail')});
    //                     }
    //                     else{
    //                         console.log('Success');
    //                         response.status(SUCCESS).json({message: message_reader.readMessage('user.updatesucess')});
    //                     }
    //                 })
    //             }
    //             else{
    //                 response.status(RESOURCE_NOT_FOUND).json({message:message_reader.readMessage('user.invalid')});
    //             }
    //         }
    //         else{
    //             // Invalid 
    //             response.status(RESOURCE_NOT_FOUND).json({message:message_reader.readMessage('user.invalid')});

    //         }
    //     })
    // },
    // add(userObject){
    //     userObject.password = encrypt.hashPassword(userObject.password);
    //     return UserModel.create(userObject);
    // },
    // find(userObject, response){
        
    //     UserModel.findOne({'email':userObject.email},(err,doc)=>{
    //         if(err){
    //             response.status(SERVER_ERROR).json({message:'Some Problem During Login '+err});
    //         }
    //         else if(doc && doc.email){
    //             let dbPassword = doc.password;
    //             let plainPassword = userObject.password;
    //             if(encrypt.compareHash(plainPassword, dbPassword)){
    //                 const token = generateToken(userObject.email);
    //                 response.status(SUCCESS).json({message:message_reader.readMessage('user.welcome')+doc.name,token : token});
    //             }
    //             else{
    //                 response.status(RESOURCE_NOT_FOUND).json({message:message_reader.readMessage('user.invalid')});
    //             }
    //         }
    //         else{
    //             // Invalid 
    //             response.status(RESOURCE_NOT_FOUND).json({message:message_reader.readMessage('user.invalid')});

    //         }
    //     })
    // },
    // update(userObject){
    //     UserModel.findOneAndUpdate({'email':userObject.email, 'password':userObject.password},{'password':'aaaaaaa'},(err,doc)=>{

    //     })
    //    // Change Password
    // },
    // remove(userObject){
    //         UserModel.findOneAndRemove({'email':userObject.email},(err)=>{

    //         })
    // }
