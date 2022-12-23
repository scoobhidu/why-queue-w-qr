import { Navigate, useNavigate } from "react-router-dom";

export const Classes = ({ cname }) => {
  const Nav=useNavigate();

  return (
    <div className="flex my-2" id="class">
      {/* <button className="btn btn-primary" style={{'width':'100%','height':'20%'}}>{cname}</button> */}
      <div className="text-2xl font-extrabold my-auto mr-2 text-purple-700">{cname}</div>
      <button
        class="shadow bg-purple-500 hover:bg-purple-400 focus:shadow-outline focus:outline-none text-white font-bold py-2 px-4 rounded mr-5 opacity-30 hover:opacity-100"
        type="button"
        onClick={()=>Nav('/qr')}
      >
        Generate QR Code
      </button>
      <button
        class="shadow bg-purple-500 hover:bg-purple-400 focus:shadow-outline focus:outline-none text-white font-bold py-2 px-4 rounded opacity-30 hover:opacity-100"
        type="button"
        onClick={()=>Nav('/attend')}
      >
        Attendance History
      </button>
    </div>
  );
};
