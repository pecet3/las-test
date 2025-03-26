import React, { ReactNode } from "react";
import { useNavigate } from "react-router-dom";
import { useAppContext } from "../contexts/AppContext";

type LogoutButtonProps = {
  children: ReactNode;
};

export const LogoutButton: React.FC<LogoutButtonProps> = ({ children }) => {
  const { setUser } = useAppContext();
  const navigate = useNavigate();

  const handleLogout = async () => {
    try {
      const response = await fetch("/api/auth/logout", {
        method: "GET",
        credentials: "include",
      });

      if (!response.ok) {
        throw new Error("Logout failed");
      }
      setUser(null);
      navigate("/");
    } catch (error) {
      console.error("Error during logout:", error);
    }
  };

  return (
    <button onClick={handleLogout} className="text-white">
      {children}
    </button>
  );
};
