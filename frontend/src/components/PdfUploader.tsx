import React, { useState, ChangeEvent, FormEvent } from "react";

type MagicLinkResponse = {
  url: string;
};

export const PdfUploader: React.FC = () => {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [uploadStatus, setUploadStatus] = useState<string>("");
  const [error, setError] = useState<string>("");

  const [magicLink, setMagicLink] = useState("");

  const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
    setUploadStatus("");
    setError("");

    if (event.target.files && event.target.files.length > 0) {
      const file = event.target.files[0];

      if (file.type !== "application/pdf") {
        setError("Only PDF files are allowed");
        return;
      }

      if (file.size > 500 * 1024) {
        setError("File size should not exceed 500kb");
        return;
      }

      setSelectedFile(file);
    }
  };

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    if (!selectedFile) {
      setError("Please select a PDF file");
      return;
    }

    const formData = new FormData();
    formData.append("file", selectedFile);

    try {
      const response = await fetch("/api/pdfs", {
        method: "POST",
        body: formData,
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(errorText || "File upload failed");
      }
      const data: MagicLinkResponse = await response.json();
      setMagicLink(data.url);
      setUploadStatus("File uploaded successfully");
      setError("");
    } catch (err) {
      setError(err instanceof Error ? err.message : "File upload failed");
      setUploadStatus("");
      console.error("Upload error:", err);
    }
  };

  return (
    <>
      <div className="max-w-md mx-auto mt-10 p-6 bg-white rounded-lg shadow-md">
        <h2 className="text-2xl font-bold mb-6 text-center">
          PDF File Uploader
        </h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <input
              type="file"
              accept=".pdf"
              onChange={handleFileChange}
              className="w-full p-2 border rounded"
            />
          </div>

          {error && <div className="mb-4 text-red-500 text-sm">{error}</div>}

          {uploadStatus && (
            <div className="mb-4 text-green-500 text-sm">{uploadStatus}</div>
          )}

          <button
            type="submit"
            disabled={!selectedFile}
            className="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600 disabled:bg-gray-400"
          >
            Upload PDF
          </button>
        </form>
      </div>
      {magicLink !== "" ? (
        <a href={magicLink} className="text-2xl text-purple-300">
          Magic Link
        </a>
      ) : null}
    </>
  );
};
