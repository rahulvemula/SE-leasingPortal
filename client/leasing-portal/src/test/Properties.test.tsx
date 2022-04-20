import Properties from '../components/Properties';
import { render, screen } from '@testing-library/react';
import { HashRouter } from 'react-router-dom';

test('renders properties', () => {
    render(<HashRouter><Properties/></HashRouter>);
    const properties = screen.getAllByTestId('properties');
    expect(properties[0]).toBeInTheDocument();

})