import React from 'react';
import { Card, Tabs } from 'antd';
import { useLocaleStore } from '../stores/useLocaleStore';

const Settings: React.FC = () => {
  const { messages } = useLocaleStore();
  
  const items = [
    {
      key: 'general',
      label: messages.settings.sections.general,
      children: <div>通用设置内容</div>,
    },
    {
      key: 'network',
      label: messages.settings.sections.network,
      children: <div>网络设置内容</div>,
    },
    {
      key: 'sync',
      label: messages.settings.sections.sync,
      children: <div>同步设置内容</div>,
    },
  ];

  return (
    <Card title={messages.settings.title}>
      <Tabs items={items} />
    </Card>
  );
};

export default Settings; 