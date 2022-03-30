import React from 'react';
import { render, screen } from '@testing-library/react';
import App from '../App';

test('renders Hous logo', () => {
  render(<App />);
  const housElement = screen.getByText('HOUS.');
  expect(housElement).toBeInTheDocument();
});

test('renders login and register buttons', () => {
  render(<App />);
  const loginElement = screen.getByText('Login');
  expect(loginElement).toBeInTheDocument();
  const registerElement = screen.getByText('Register');
  expect(registerElement).toBeInTheDocument();
});
