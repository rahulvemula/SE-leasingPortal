import LeaseConfirm from "../pages/lease-confirmation";
import { render, screen } from "@testing-library/react";
import { BrowserRouter } from "react-router-dom";
import { Provider } from "react-redux";
import configureStore from "redux-mock-store";

test("shows all lease details", () => {
  const initialState = { leaseData: { name: "lahari" } };
  const mockStore = configureStore();
  let store = mockStore(initialState);

  render(
    <Provider store={store}>
      <BrowserRouter>
        <LeaseConfirm />
      </BrowserRouter>
    </Provider>
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
