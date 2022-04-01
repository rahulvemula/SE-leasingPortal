import LeaseConfirm from "../pages/lease-confirmation";
import { render, screen } from "@testing-library/react";
import { BrowserRouter } from "react-router-dom";

test("shows all lease details", () => {
  render(
    <BrowserRouter>
      <LeaseConfirm />
    </BrowserRouter>
  );

  const nameElement = screen.getByText("Name");
  expect(nameElement).toBeInTheDocument();

  const emailElement = screen.getByText("Email");
  expect(emailElement).toBeInTheDocument();

  const pNameElement = screen.getByText("Property Name");
  expect(pNameElement).toBeInTheDocument();

  const startDateElement = screen.getByText("Lease start date");
  expect(startDateElement).toBeInTheDocument();

  const endDateElement = screen.getByText("Lease end date");
  expect(endDateElement).toBeInTheDocument();
});
