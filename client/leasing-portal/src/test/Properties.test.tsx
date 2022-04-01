import Properties from '../components/Properties';
import { render, screen } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';

test('renders properties', () => {
    render(<BrowserRouter><Properties/></BrowserRouter>);
    const properties = screen.getAllByTestId('properties');
    expect(properties[0]).toBeInTheDocument();

})