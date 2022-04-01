import LeaseConfirm from '../pages/lease-confirmation';
import { render, screen } from '@testing-library/react';

test('shows all lease details', () => {
    render(<LeaseConfirm/>);

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