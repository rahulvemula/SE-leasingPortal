import Listing from '../pages/Listing';
import { render, screen } from '@testing-library/react';
import { BrowserRouter } from "react-router-dom";


test('render lisiting page', () =>  {
    const testComponent = render(
    <BrowserRouter>
        <Listing/>
    </BrowserRouter>
    );
        

    const nameElement = screen.getByText("Name");
    expect(nameElement).toBeInTheDocument();

    const nameInput = testComponent.container.querySelector('#registerInputName');
    expect(nameInput).toBeInTheDocument();

    const emailElement = screen.getByText("Email address");
    expect(emailElement).toBeInTheDocument();

    const emailInput = testComponent.container.querySelector('#registerInputEmail');
    expect(emailInput).toBeInTheDocument();

    const startDateElement = screen.getByText("Start date");
    expect(startDateElement).toBeInTheDocument();

    const startDateInput = testComponent.container.querySelector('#registerInputStartDate');
    expect(startDateInput).toBeInTheDocument();

    const endDateElement = screen.getByText("End date");
    expect(endDateElement).toBeInTheDocument();

    const endDateInput = testComponent.container.querySelector('#registerInputStartDate');
    expect(endDateInput).toBeInTheDocument();
});