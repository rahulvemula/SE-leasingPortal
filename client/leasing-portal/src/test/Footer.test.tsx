import React from 'react';
import { render, screen } from '@testing-library/react';
import Footer from '../components/Footer';

test('renders Hous logo', () => {
  render(<Footer />);
  const housElement = screen.getByText('Hous.');
  expect(housElement).toBeInTheDocument();
});

test('renders community heading', () => {
  render(<Footer />);
  const communityElement = screen.getByText('Community');
  expect(communityElement).toBeInTheDocument();
});

test('renders About us links', () => {
  render(<Footer />);
  const aboutUsElement = screen.getByText('Our Story');
  expect(aboutUsElement).toBeInTheDocument();
  const teamElement = screen.getByText('Meet the team');
  expect(teamElement).toBeInTheDocument();
});
test('renders company heading', () => {
  render(<Footer />);
  const companyElement = screen.getByText('Company');
  expect(companyElement).toBeInTheDocument();
});
test('renders Questions heading', () => {
  render(<Footer />);
  const qElement = screen.getByText('Have Questions?');
  expect(qElement).toBeInTheDocument();
});