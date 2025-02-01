import { useState } from 'react';
import { Layout, Menu, Button, theme } from 'antd';
import {
  DashboardOutlined,
  BlockOutlined,
  TransactionOutlined,
  SettingOutlined,
  TranslationOutlined
} from '@ant-design/icons';
import { Outlet, useNavigate } from 'react-router-dom';
import { useLocaleStore } from '../stores/useLocaleStore';

const { Header, Sider, Content } = Layout;

const MainLayout = () => {
  const [collapsed, setCollapsed] = useState(false);
  const navigate = useNavigate();
  const { token } = theme.useToken();
  const { locale, messages, setLocale } = useLocaleStore();

  const toggleLocale = () => {
    setLocale(locale === 'zh_CN' ? 'en_US' : 'zh_CN');
  };

  const menuItems = [
    {
      key: 'dashboard',
      icon: <DashboardOutlined />,
      label: messages.menu.dashboard,
    },
    {
      key: 'blocks',
      icon: <BlockOutlined />,
      label: messages.menu.blocks,
    },
    {
      key: 'transactions',
      icon: <TransactionOutlined />,
      label: messages.menu.transactions,
    },
    {
      key: 'settings',
      icon: <SettingOutlined />,
      label: messages.menu.settings,
    },
  ];

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider collapsible collapsed={collapsed} onCollapse={setCollapsed}>
        <div
          style={{
            height: 64,
            margin: 16,
            background: 'rgba(255, 255, 255, 0.1)',
            borderRadius: 6,
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            color: token.colorWhite,
            fontSize: collapsed ? 24 : 18,
            overflow: 'hidden',
            whiteSpace: 'nowrap',
          }}
        >
          <BlockOutlined style={{ marginRight: collapsed ? 0 : 8 }} />
          {!collapsed && <span>{messages.common.platformName}</span>}
        </div>
        <Menu
          theme="dark"
          defaultSelectedKeys={['dashboard']}
          mode="inline"
          items={menuItems}
          onClick={({ key }) => navigate(`/${key}`)}
        />
      </Sider>
      <Layout>
        <Header style={{ 
          padding: '0 16px', 
          background: '#fff',
          display: 'flex',
          justifyContent: 'flex-end',
          alignItems: 'center'
        }}>
          <Button
            icon={<TranslationOutlined />}
            onClick={toggleLocale}
          >
            {messages.header.switchLang}
          </Button>
        </Header>
        <Content style={{ margin: '16px' }}>
          <Outlet />
        </Content>
      </Layout>
    </Layout>
  );
};

export default MainLayout; 