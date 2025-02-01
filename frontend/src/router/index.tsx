import { RouteObject, Navigate } from 'react-router-dom';
import Dashboard from '../pages/Dashboard';
import Blocks from '../pages/Blocks';
import Transactions from '../pages/Transactions';
import Settings from '../pages/Settings';

export const routes: RouteObject[] = [
  {
    path: '/',
    element: <Navigate to="/dashboard" replace />,
  },
  {
    path: '/dashboard',
    element: <Dashboard />,
  },
  {
    path: '/blocks',
    element: <Blocks />,
  },
  {
    path: '/transactions',
    element: <Transactions />,
  },
  {
    path: '/settings',
    element: <Settings />,
  },
]; 