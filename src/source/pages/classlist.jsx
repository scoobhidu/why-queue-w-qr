import { useSelector } from "react-redux";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Classes } from "../components/classes";
import { getClass } from "../utils/axios/getclass";
// import Logout from "../components/logout";
import logoutcall from "../utils/logoutcall";

export const ClassList = () => {
  const state = useSelector((state) => state);
  const classofteacher = state.classes;
  // const [loggedin, setloggedin] = useState(state.loggedin)
  const loggedin = state.loggedin;
  const token = state.token;
  console.log(classofteacher);

  const Navigate = useNavigate();

  const handleLogout = (event) => {
    // event.preventDefault();
    console.log("hello");
    logoutcall();
  };

  console.log(loggedin)

  if (loggedin === false) {
    Navigate("/")
  } else if (loggedin && classofteacher.length === 0) {
    //can find better condition
    console.log("calling getclass");
    getClass(token);
  } else {
    console.log("printing classlist");
    return (
      <div>
        <div className="flex justify-end">
        <button onClick={handleLogout} className="shadow text-sm bg-purple-500  hover:bg-purple-400 focus:shadow-outline focus:outline-none text-white font-medium py-2 px-4 rounded m-2">Logout</button>
        </div>
        
        <div className=" py-10 ">
          <div className="bg-white  w-4/5 m-auto rounded-3xl opacity-100 py-5">
            <div className="text-4xl font-bold mb-5 pt-2 text-purple-800">
              Class List 
            </div>
            {classofteacher.map((element) => {
              return (
                <div>
                  <div className="flex justify-center">
                    <Classes cname={element} key={element} />
                  </div>
                  <div></div>
                </div>
              );
            })}
          </div>
        </div>
      </div>
    );
  }
};
