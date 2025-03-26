import { useAppContext } from "../contexts/AppContext";

export const Navbar = () => {
  const { user } = useAppContext();
  return (
    <>
      <nav className="flex justify-between items-center w-full px-1 sm:px-4 pt-1"></nav>
    </>
  );
};
