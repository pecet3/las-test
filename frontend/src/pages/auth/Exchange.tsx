import React, { useEffect, useState } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";

const API_URL = "/api";
const PREFIX = "/auth";

export const Exchange: React.FC = () => {
  const [searchParams] = useSearchParams();
  const email = searchParams.get("email");
  const name = searchParams.get("name");
  const type = searchParams.get("type");
  const navigate = useNavigate();

  const [code, setCode] = useState("");
  const [loading, setLoading] = useState(false);
  const [resendDisabled, setResendDisabled] = useState(false);
  const [timer, setTimer] = useState(30);

  useEffect(() => {
    let interval: any;
    if (resendDisabled) {
      interval = setInterval(() => {
        setTimer((prev) => {
          if (prev <= 1) {
            clearInterval(interval);
            setResendDisabled(false);
            return 30;
          }
          return prev - 1;
        });
      }, 1000);
    }
    return () => clearInterval(interval);
  }, [resendDisabled]);

  const handleResend = async () => {
    if (resendDisabled || !email || !type) return;
    setResendDisabled(true);
    setLoading(true);

    try {
      const endpoint = type === "register" ? "register" : "login";
      const body =
        type === "register"
          ? JSON.stringify({ email, name })
          : JSON.stringify({ email });
      const response = await fetch(`${API_URL}${PREFIX}/${endpoint}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body,
      });

      if (response.ok) {
        alert("Verification code resent!");
      } else {
        const message = await response.text();
        alert(message || "Something went wrong...");
      }
    } catch (error) {
      alert("Something went wrong...");
    }
    setLoading(false);
  };

  const handleExchange = async () => {
    if (code.length !== 6 || !email || !type) return;
    setLoading(true);

    try {
      const endpoint =
        type === "register" ? "register/exchange" : "login/exchange";
      const response = await fetch(`${API_URL}${PREFIX}/${endpoint}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, code }),
      });

      if (response.ok) {
        navigate("/");
      } else {
        const message = await response.text();
        alert(message || "Something went wrong...");
      }
    } catch (error) {
      alert("Something went wrong...");
    }
    setLoading(false);
  };

  return (
    <main className="min-h-[68vh] flex flex-col items-center gap-2 justify-center h-screen pb-64 p-1">
      <h1 className="text-2xl font-bold">Verify Your Email</h1>
      <p className="text-xl text-center">
        Please check your email for the verification code.
      </p>

      <div className="w-full max-w-md bg-white border border-black p-4 rounded-md shadow-md flex flex-col gap-2">
        <input
          className="bg-gray-200 w-full border border-black rounded-md p-2 font-mono"
          type="text"
          placeholder="Code"
          value={code}
          onChange={(e) => setCode(e.target.value)}
          maxLength={6}
        />
        <button
          className={`w-full btn bg-teal-200 text-white p-3 rounded-md ${
            loading ? "opacity-50" : ""
          }`}
          onClick={handleExchange}
          disabled={loading}
        >
          {loading ? "Loading..." : "Submit"}
        </button>
        <button
          className="text-center w-full text-blue-600"
          onClick={handleResend}
          disabled={resendDisabled || loading}
        >
          {resendDisabled ? `Try again in ${timer}s` : "Resend an email"}
        </button>
      </div>
    </main>
  );
};
