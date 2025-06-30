import React, { useState } from "react";
import { SIGNUP_URL } from "../const/const";

const Signup = () => {
  const [username, setUsername] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    const response = await fetch(SIGNUP_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ username }),
    });

    if (response.ok) {
      alert("Signup successful");
      setUsername("");
    } else {
      alert("Signup failed");
    }
  };

  return (
    <div style={styles.container}>
      <form style={styles.form} onSubmit={handleSubmit}>
        <h2 style={styles.heading}>Signup Form</h2>

        <input
          type="text"
          placeholder="Enter username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
          style={styles.input}
        />

        <button type="submit" style={styles.button}>
          Submit
        </button>
      </form>
    </div>
  );
};

const styles = {
  container: {
    display: "flex",
    height: "100vh",
    justifyContent: "center",
    alignItems: "center",
    background: "#f0f4f8",
  },
  form: {
    padding: "30px",
    borderRadius: "10px",
    backgroundColor: "#ffffff",
    boxShadow: "0px 4px 10px rgba(0,0,0,0.1)",
    display: "flex",
    flexDirection: "column",
    width: "300px",
  },
  heading: {
    marginBottom: "20px",
    textAlign: "center",
    color: "#333",
  },
  input: {
    padding: "10px",
    fontSize: "16px",
    marginBottom: "15px",
    borderRadius: "5px",
    border: "1px solid #ccc",
  },
  button: {
    padding: "10px",
    backgroundColor: "#4CAF50",
    color: "white",
    fontSize: "16px",
    border: "none",
    borderRadius: "5px",
    cursor: "pointer",
  },
};

export default Signup;
