import React, { useEffect, useState } from "react";
import { Loading } from "./Loading";

interface PDF {
  uuid: string;
  name: string;
  created_at: string;
  last_open_at: string;
  url: string;
}

interface PDFViewerProps {
  url: string;
  onClose: () => void;
}

const PDFViewer: React.FC<PDFViewerProps> = ({ url, onClose }) => {
  return (
    <div
      className="fixed inset-0 z-50 bg-white bg-opacity-20
    flex items-center justify-center"
    >
      <div className="relative w-5/6 h-5/6 bg-white rounded-lg shadow-xl">
        <button
          onClick={onClose}
          className="absolute top-4 right-4 bg-red-500 text-white px-4 py-2 rounded"
        >
          Close
        </button>
        <iframe src={url} className="w-full h-full" title="PDF Viewer" />
      </div>
    </div>
  );
};

interface PDFListProps {
  pdfs: PDF[];
  setPdfs: React.Dispatch<React.SetStateAction<PDF[]>>;
}

const PDFList: React.FC<PDFListProps> = ({ pdfs, setPdfs }) => {
  const [selectedPdfUrl, setSelectedPdfUrl] = useState<string | null>(null);

  const onDeletePdf = async (uuid: string) => {
    try {
      const response = await fetch(`/api/pdfs/${uuid}`, {
        method: "DELETE",
      });

      if (!response.ok) {
        throw new Error(`Error deleting PDF: ${response.statusText}`);
      }

      setPdfs((prevPdfs) => prevPdfs.filter((pdf) => pdf.uuid !== uuid));
    } catch (error) {
      alert("Failed to delete PDF");
    }
  };

  return (
    <div>
      {selectedPdfUrl && (
        <PDFViewer
          url={selectedPdfUrl}
          onClose={() => setSelectedPdfUrl(null)}
        />
      )}

      <h2 className="text-2xl font-bold mb-6">PDF Files</h2>

      {pdfs.length === 0 ? (
        <p>No PDF files found.</p>
      ) : (
        <ul
          className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 
          m-auto w-full gap-4 sm:gap-4 md:gap-6 max-w-4xl"
        >
          {pdfs.map((pdf) => (
            <li
              key={pdf.uuid}
              className={`cursor-pointer w-full max-w-md border border-black 
              py-4 px-6 rounded-md shadow-md flex flex-col m-auto h-full 
              hover:scale-105 duration-500 hover:shadow-xl bg-white`}
            >
              <div
                className={`flex flex-col items-start justify-between 
                m-auto h-full w-full`}
              >
                <h3 className="text-base font-semibold">{pdf.name}</h3>

                <p className="self-end text-xs text-gray-500">
                  Last open at: {new Date(pdf.last_open_at).toLocaleString()}
                </p>

                <p className="self-end text-xs text-gray-500">
                  Created at: {new Date(pdf.created_at).toLocaleString()}
                </p>

                <span className="flex items-center m-auto justify-between w-full">
                  <button
                    className="bg-blue-500 text-white px-3 py-1 rounded"
                    onClick={() => setSelectedPdfUrl(pdf.url)}
                  >
                    View PDF
                  </button>

                  <button
                    className="bg-red-500 text-white px-3 py-1 rounded"
                    onClick={() => onDeletePdf(pdf.uuid)}
                  >
                    Delete
                  </button>
                </span>
              </div>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export const Explorer = () => {
  const [pdfData, setPdfData] = useState<PDF[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchPdfs = async () => {
      try {
        const response = await fetch("/api/pdfs");

        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data: PDF[] = await response.json();

        if (data) {
          setPdfData(data);
        }
      } catch (e: any) {
        setError(e.message || "An error occurred while fetching PDFs.");
      } finally {
        setLoading(false);
      }
    };

    fetchPdfs();
  }, []);

  if (loading) {
    return <Loading />;
  }

  if (error) {
    return <p>Error: {error}</p>;
  }

  return (
    <div>
      {pdfData.length === 0 ? (
        <p className="my-24 text-3xl">No PDFs...</p>
      ) : (
        <PDFList pdfs={pdfData} setPdfs={setPdfData} />
      )}
    </div>
  );
};
