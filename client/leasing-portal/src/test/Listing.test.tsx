import Listing from "../pages/Listing";
import { render, screen } from "@testing-library/react";
import { HashRouter } from "react-router-dom";
import { Provider } from "react-redux";
import configureStore from "redux-mock-store";

test("render lisiting page", () => {
  const initialState = { leaseData: { name: "lahari" } };
  const mockStore = configureStore();
  let store = mockStore(initialState);
  const testComponent = render(
    <Provider store={store}>
      <HashRouter>
        <Listing />
      </HashRouter>
    </Provider>
  );

  const nameElement = screen.getByText("Name");
  expect(nameElement).toBeInTheDocument();

  const nameInput = testComponent.container.querySelector("#registerInputName");
  expect(nameInput).toBeInTheDocument();

  const emailElement = screen.getByText("Email address");
  expect(emailElement).toBeInTheDocument();

  const emailInput = testComponent.container.querySelector(
    "#registerInputEmail"
  );
  expect(emailInput).toBeInTheDocument();

  const startDateElement = screen.getByText("Start date");
  expect(startDateElement).toBeInTheDocument();

  const startDateInput = testComponent.container.querySelector(
    "#registerInputStartDate"
  );
  expect(startDateInput).toBeInTheDocument();

  const endDateElement = screen.getByText("End date");
  expect(endDateElement).toBeInTheDocument();

  const endDateInput = testComponent.container.querySelector(
    "#registerInputStartDate"
  );
  expect(endDateInput).toBeInTheDocument();
});
