import "./App.css";
import { Route, Routes } from "react-router-dom";
import { ProtectedPage } from "./wrappers/Protected";
import { Home } from "./pages/Home";
import { Navbar } from "./components/Navbar";
import { Exchange } from "./pages/auth/Exchange";
import { Auth } from "./pages/auth/Auth";
import { Upload } from "./pages/Upload";
import { Explorer } from "./components/Explorer";

function App() {
  return (
    <>
      <Routes>
        <Route
          path="/"
          element={
            <ProtectedPage>
              <Navbar />
              <Home />
            </ProtectedPage>
          }
        />
        <Route
          path="/upload"
          element={
            <ProtectedPage>
              <Navbar />
              <Upload />
            </ProtectedPage>
          }
        />
        <Route
          path="/explorer"
          element={
            <ProtectedPage>
              <Navbar />
              <Explorer />
            </ProtectedPage>
          }
        />
        <Route
          path="/auth"
          element={
            <>
              <Navbar />
              <Auth />
            </>
          }
        />
        <Route
          path="/auth/exchange"
          element={
            <>
              <Navbar />
              <Exchange />
            </>
          }
        />
      </Routes>
    </>
  );
}

export default App;
