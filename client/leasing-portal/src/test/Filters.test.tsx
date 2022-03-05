import React from 'react';
import { fireEvent, render, screen } from '@testing-library/react';
import Filters from '../components/Filters';

test('renders with out crashing', () => {
  render(<Filters />);
});

test('locations options diplayed', () => {
  render(<Filters/>);
  const dropdown = screen.getAllByTestId('locationsSelect');
  expect(dropdown[0].children[0].textContent).toBe('Gainesville');
  expect(dropdown[0].children[1].textContent).toBe('Santa Fe');
  expect(dropdown[0].children[2].textContent).toBe('Alachua');
});

test('property types options diplayed', () => {
  render(<Filters/>);
  const dropdown = screen.getAllByTestId('propTypesSelect');
  expect(dropdown[0].children[0].textContent).toBe('Residential');
  expect(dropdown[0].children[1].textContent).toBe('Commercial');
  expect(dropdown[0].children[2].textContent).toBe('Office');
});

test('prices options diplayed', () => {
  render(<Filters/>);
  const dropdown = screen.getAllByTestId('pricesSelect');
  expect(dropdown[0].children[0].textContent).toBe('250');
  expect(dropdown[0].children[3].textContent).toBe('1000');
  expect(dropdown[0].children[5].textContent).toBe('5000');
});
