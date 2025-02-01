import { ConfigProvider } from 'antd';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import Layout from './layouts/MainLayout';
import Dashboard from './pages/Dashboard';
import Blocks from './pages/Blocks';
import Transactions from './pages/Transactions';
import Settings from './pages/Settings';
import { useLocaleStore } from './stores/useLocaleStore';

const queryClient = new QueryClient();

function App() {
  const { antdLocale } = useLocaleStore();

  return (
    <QueryClientProvider client={queryClient}>
      <ConfigProvider locale={antdLocale}>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Layout />}>
              <Route index element={<Navigate to="/dashboard" replace />} />
              <Route path="dashboard" element={<Dashboard />} />
              <Route path="blocks" element={<Blocks />} />
              <Route path="transactions" element={<Transactions />} />
              <Route path="settings" element={<Settings />} />
            </Route>
          </Routes>
        </BrowserRouter>
      </ConfigProvider>
    </QueryClientProvider>
  );
}

export default App; 