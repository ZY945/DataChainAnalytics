import React from 'react';
import { Card } from 'antd';
import { useLocaleStore } from '../stores/useLocaleStore';

const Dashboard: React.FC = () => {
  const { messages } = useLocaleStore();
  
  return (
    <Card title={messages.menu.dashboard}>
      <p>{messages.common.welcome}</p>
    </Card>
  );
};

export default Dashboard; 