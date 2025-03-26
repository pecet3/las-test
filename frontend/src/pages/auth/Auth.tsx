import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const API_URL = "/api";
const PREFIX = "/auth";

type Step = "register" | "login";

interface FormData {
  name: string;
  email: string;
}
const nameValidate = (name: string): boolean => {
  if (name.length < 2 || name.length > 16) {
    alert("Name must be between 2 and 16 characters.");
    return false;
  }
  return true;
};

export const Auth: React.FC = () => {
  const navigate = useNavigate();

  const [currentStep, setCurrentStep] = useState<Step>("register");
  const [formData, setFormData] = useState<FormData>({
    name: "",
    email: "",
  });
  const [loading, setLoading] = useState<boolean>(false);
  const [isChecked, setIsChecked] = useState(false);

  const handleRegister = async () => {
    if (!isChecked) {
      alert("You have to accept Terms & Conditions");
      return;
    }
    if (formData.email.length < 5 || !nameValidate(formData.name)) {
      alert("Email or name doesn't meet our criteria");
      return;
    }
    setLoading(true);
    try {
      const response = await fetch(`${API_URL}${PREFIX}/register`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: formData.name,
          email: formData.email,
        }),
      });

      if (response.status === 200) {
        navigate(
          `/auth/exchange?email=${formData.email}&name=${formData.name}&type=register`
        );
      } else {
        const message = await response.text();
        if (message && message.length > 2) {
          alert(message);
        } else {
          alert("Something went wrong...");
        }
      }
    } catch (error) {
      alert("Something went wrong...");
    }
    setLoading(false);
  };

  const handleLogin = async () => {
    if (formData.email.length < 5) {
      alert("Provided wrong email");
      return;
    }
    setLoading(true);
    try {
      const response = await fetch(`${API_URL}${PREFIX}/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email: formData.email,
        }),
      });

      if (response.ok) {
        navigate(`/auth/exchange?email=${formData.email}&type=login`);
      } else {
        const message = await response.text();
        if (message && message.length > 2) {
          alert(message);
        } else {
          alert("Something went wrong...");
        }
      }
    } catch (error) {
      alert("Something went wrong...");
    }
    setLoading(false);
  };

  const renderRegisterForm = () => (
    <>
      <input
        className="bg-gray-200 w-full border border-black rounded-md p-2 font-mono"
        type="text"
        placeholder="Name"
        value={formData.name}
        onChange={(e) => setFormData({ ...formData, name: e.target.value })}
      />
      <input
        className="bg-gray-200 w-full border border-black rounded-md p-2 font-mono"
        type="email"
        placeholder="Email"
        value={formData.email}
        onChange={(e) => setFormData({ ...formData, email: e.target.value })}
      />
      <div className="flex gap-2 mt-2 px-1">
        <input
          type="checkbox"
          checked={isChecked}
          onChange={(e) => setIsChecked(e.target.checked)}
        />
        <p className="text-sm">
          I have read and accept the{" "}
          <b className="underline text-blue-600">Terms & Conditions</b>.
        </p>
      </div>
      <button
        className={`w-full btn bg-teal-200 text-white p-3 rounded-md mb-4 ${
          loading ? "opacity-50" : ""
        }`}
        onClick={handleRegister}
        disabled={loading}
      >
        {loading ? "Loading..." : "Submit"}
      </button>
      <p
        className="text-center text-blue-600 cursor-pointer"
        onClick={() => setCurrentStep("login")}
      >
        Already have an account? Sign in
      </p>
    </>
  );

  const renderLoginForm = () => (
    <>
      <input
        className="bg-gray-200 w-full border border-black rounded-md p-2 font-mono"
        type="email"
        placeholder="Email"
        value={formData.email}
        onChange={(e) => setFormData({ ...formData, email: e.target.value })}
      />
      <button
        className={`w-full btn bg-teal-200 text-white p-3 rounded-md mb-4 ${
          loading ? "opacity-50" : ""
        }`}
        onClick={handleLogin}
        disabled={loading}
      >
        {loading ? "Loading..." : "Submit"}
      </button>
      <p
        className="text-center text-blue-600 cursor-pointer"
        onClick={() => setCurrentStep("register")}
      >
        Don't have an account? Sign up
      </p>
    </>
  );

  return (
    <main className="min-h-[68vh] flex flex-col items-center gap-2 justify-center h-screen pb-64 p-1">
      <h1 className="text-2xl font-bold">
        {currentStep === "register" ? "Sign Up" : "Sign In"}
      </h1>
      <div className="w-96  bg-white border border-black p-4 rounded-md shadow-md flex flex-col gap-2">
        {currentStep === "register" && renderRegisterForm()}
        {currentStep === "login" && renderLoginForm()}
      </div>
    </main>
  );
};
