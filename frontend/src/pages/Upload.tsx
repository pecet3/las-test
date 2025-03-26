import { PdfUploader } from "../components/PdfUploader";

export const Upload = () => {
  return (
    <main className="min-h-[68vh] flex flex-col gap-10 sm:gap-16 h-full">
      <PdfUploader />
    </main>
  );
};
