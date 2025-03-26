import { Link } from "react-router-dom";
import { LogoutButton } from "./Logout";

export const Navbar = () => {
  return (
    <nav className="flex justify-between items-center w-full px-1 sm:px-4 pt-1">
      <div className="flex items-center gap-1">
        <Link to={"/"} className="font-mono text-2xl">
          PDF-LAS
        </Link>
      </div>
      <Link to={"/upload"} className="">
        <button className="text-white">Upload</button>
      </Link>
      <div className="justify-end items-center flex">
        <LogoutButton>Logout</LogoutButton>
      </div>
    </nav>
  );
};
