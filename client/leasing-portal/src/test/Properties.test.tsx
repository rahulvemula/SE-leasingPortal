import Properties from '../components/Properties';
import { render, screen } from '@testing-library/react';

test('renders properties', () => {
    render(<Properties/>);
    const properties = screen.getAllByTestId('properties');
    expect(properties[0]).toBeInTheDocument();

})