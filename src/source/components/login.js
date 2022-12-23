import { useEffect, useState } from "react";
import { Navigate } from "react-router-dom";
import { LoginCall } from "../utils/axios/login";

export const LoginValidation = () => {
  function verify() {
    if (userName !== "" && ID !== "") {
      console.log("Send Data");
      LoginCall(userName, ID);
    }
  }

  const [userName, setUser] = useState("");
  const [ID, setID] = useState("");

//   useEffect(() => {
//     if (userName === "" || ID === "") {
//       document.getElementById("form").className = "btn btn-danger disabled";
//     } else {
//       document.getElementById("form").className = "btn btn-success";
//     }
//   }, [userName, ID]);

  return (
    <div className="container bg-white rounded-3xl my-20">
      <div className="text-4xl pt-10 font-bold text-purple-800">Login to the Portal</div>
      <form className="py-16">
        <div className="">
          {/* <label htmlFor="userName" className="">
            UserName
          </label>{" "} */}
          &nbsp;
          <input
            type="text"
            className="bg-gray-200 appearance-none border-2 border-gray-200 rounded py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500 w-1/2"
            id="userName"
            placeholder="Username"
            onChange={(e) => setUser(e.target.value)}
          />
        </div>
        <br />
        <div className="">
          {/* <label htmlFor="id" className="display-6 form-label">
            Special ID
          </label>{" "} */}
          &nbsp;
          <input
            type="password"
            className="bg-gray-200 appearance-none border-2 border-gray-200 rounded py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500 w-1/2"
            id="id"
            placeholder="password"
            onChange={(e) => setID(e.target.value)}
          />
        </div>
        <div>
          <button
            id="form"
            className="shadow bg-purple-500  hover:bg-purple-400 focus:shadow-outline focus:outline-none text-white font-bold py-2 px-4 rounded mt-8"
            type="button"
            onClick={() => verify()}
          >
            Login Now
          </button>
          
        </div>
        
      </form>
    </div>
  );
};
