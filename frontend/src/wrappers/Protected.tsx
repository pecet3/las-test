import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAppContext } from "../contexts/AppContext";
import { Loading } from "../components/Loading";

const max_counter = 3;

export const ProtectedPage = ({ children }: { children: React.ReactNode }) => {
  const { user, setUser } = useAppContext();
  const navigate = useNavigate();
  let counter = 0;
  const fetchUser = async () => {
    try {
      const result = await fetch("/api/auth/ping");
      const data = await result.json();
      if (result.ok) {
        setUser(data);
        return;
      }
    } catch (err: any) {
      if (counter < max_counter) {
        counter++;
        await new Promise((resolve) => setTimeout(resolve, 100));
        fetchUser();
      } else {
        navigate("/auth");
      }
    }
  };

  useEffect(() => {
    if (!user) {
      fetchUser();
    }
  }, []);
  return (
    <>
      {user ? (
        <>{children}</>
      ) : (
        <div className="my-40">
          <Loading />
        </div>
      )}
    </>
  );
};
