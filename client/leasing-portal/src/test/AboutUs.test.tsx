import React from "react";
import { render, screen } from "@testing-library/react";
import AboutUs from "../components/AboutUs/AboutUs";
import { HashRouter } from "react-router-dom";

test("renders rahul name", () => {
  render(
      <AboutUs />
  );

  const rahulElement = screen.getByText("Rahul Vemula");
  expect(rahulElement).toBeInTheDocument();
});

test("renders lahari name", () => {
  render(
      <AboutUs />
  );

  const lahariElement = screen.getByText("Lahari Barad");
  expect(lahariElement).toBeInTheDocument();
});

test("renders mitul name", () => {
  render(
      <AboutUs />
  );

  const mElement = screen.getByText("Mitul Mandaliya");
  expect(mElement).toBeInTheDocument();
});

test("renders vamsi name", () => {
  render(
      <AboutUs />
  );

  const vElement = screen.getByText("Vamsi Viswanath");
  expect(vElement).toBeInTheDocument();
});
