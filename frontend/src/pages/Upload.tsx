import { PdfUploader } from "../components/PdfUploader";
import { useAppContext } from "../contexts/AppContext";

export const Upload = () => {
  const { user } = useAppContext();
  return (
    <main className="min-h-[68vh] flex flex-col gap-10 sm:gap-16 h-full">
      <PdfUploader />
    </main>
  );
};
