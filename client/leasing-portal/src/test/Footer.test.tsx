import React from "react";
import { render, screen } from "@testing-library/react";
import Footer from "../components/Footer";
import { BrowserRouter } from "react-router-dom";

test("renders Hous logo", () => {
  render(
    <BrowserRouter>
      <Footer />
    </BrowserRouter>
  );

  const housElement = screen.getByText("Hous.");
  expect(housElement).toBeInTheDocument();
});

test("renders community heading", () => {
  render(
    <BrowserRouter>
      <Footer />
    </BrowserRouter>
  );

  const communityElement = screen.getByText("Community");
  expect(communityElement).toBeInTheDocument();
});

test("renders About us links", () => {
  render(
    <BrowserRouter>
      <Footer />
    </BrowserRouter>
  );
  const aboutUsElement = screen.getByText("Our Story");
  expect(aboutUsElement).toBeInTheDocument();
  // const teamElement = screen.getByText('About Us');
  // expect(teamElement).toBeInTheDocument();
});
test("renders company heading", () => {
  render(
    <BrowserRouter>
      <Footer />
    </BrowserRouter>
  );

  const companyElement = screen.getByText("Company");
  expect(companyElement).toBeInTheDocument();
});
test("renders Questions heading", () => {
  render(
    <BrowserRouter>
      <Footer />
    </BrowserRouter>
  );

  const qElement = screen.getByText("Have Questions?");
  expect(qElement).toBeInTheDocument();
});
