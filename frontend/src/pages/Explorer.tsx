import React, { useEffect, useState } from "react";

interface PDF {
  uuid: string;
  name: string;
  createdAt: string; // Assuming the date is coming as a string
  lastOpenAt: string; // Assuming the date is coming as a string
  url: string;
}

interface PDFListProps {
  pdfs: PDF[];
}

const PDFList: React.FC<PDFListProps> = ({ pdfs }) => {
  return (
    <div>
      <h2>PDF Files</h2>
      {pdfs.length === 0 ? (
        <p>No PDF files found.</p>
      ) : (
        <ul>
          {pdfs.map((pdf) => (
            <li key={pdf.uuid}>
              <a href={pdf.url} target="_blank" rel="noopener noreferrer">
                {pdf.name}
              </a>
              <p>Created At: {new Date(pdf.createdAt).toLocaleDateString()}</p>
              <p>
                Last Opened At: {new Date(pdf.lastOpenAt).toLocaleDateString()}
              </p>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

interface AppProps {
  // ... other props for your app (if any)
}

export const Explorer: React.FC<AppProps> = (props) => {
  const [pdfData, setPdfData] = useState<PDF[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchPdfs = async () => {
      try {
        const response = await fetch("/api/pdfs"); // Replace with your actual API endpoint
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data: PDF[] = await response.json();
        setPdfData(data);
      } catch (e: any) {
        setError(e.message || "An error occurred while fetching PDFs.");
      } finally {
        setLoading(false);
      }
    };

    fetchPdfs();
  }, []);

  if (loading) {
    return <p>Loading PDF data...</p>;
  }

  if (error) {
    return <p>Error: {error}</p>;
  }

  return (
    <div>
      {/* ... other parts of your app */}
      <PDFList pdfs={pdfData} />
      {/* ... other parts of your app */}
    </div>
  );
};
