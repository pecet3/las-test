import { useAppContext } from "../contexts/AppContext";

export const Home = () => {
  const { user } = useAppContext();
  return (
    <main className="min-h-[68vh] flex flex-col gap-10 sm:gap-16 h-full">
      hello {user?.name}
    </main>
  );
};
