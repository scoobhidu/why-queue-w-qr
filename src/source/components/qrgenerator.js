import axios from "axios";
import { useState } from "react";
import { useSelector } from "react-redux";
import { Navigate } from "react-router-dom";
import { qrcall } from "../utils/axios/qrcall";
// var resizebase64 = require('resize-base64');

export const QR = () => {
  // var  img = resizebase64(QR, 100, 100);

  const loggedin = useSelector((state) => state.loggedin);
  const QR = useSelector((state) => state.QR);
  const token = useSelector((state) => state.token);

  function handleNavClasses(){
    Navigate('/classes')
  }

  if (!loggedin) {
    Navigate("/");
  } else if (QR === "") {
    qrcall(token);
  }


  // const handleNavClasses = (event) =>{
  //   event.preventDefault()
  //   Navigate("/classes")
  // }

  return (
    <>
    <div className="flex justify-start">
    <button className="shadow bg-purple-500  hover:bg-purple-400 focus:shadow-outline focus:outline-none text-white font-medium text-sm py-2 px-4 rounded mt-2 ml-2" onClick={()=>handleNavClasses()}>&larr; Back to ClassList</button>
    </div>
      <div className="w-full h-screen grid content-center border-black">
        <div className="text-2xl font-semibold pb-10">
          Scan the QR Code to Mark the Attendance
        </div>
        <div className="w-full md:w-1/2 lg:w-1/4 m-auto">
          <img
            src={QR}
            alt=""
            className="w-11/12 m-auto border-black border-2 rounded-2xl"
          />
        </div>
        <button
          onClick={() => qrcall(token)}
          className="shadow bg-purple-500  hover:bg-purple-400 focus:shadow-outline focus:outline-none text-white font-bold py-2 px-4 rounded mt-8 mx-auto"
        >
          Refresh the QR
        </button>
      </div>
    </>
  );
};
