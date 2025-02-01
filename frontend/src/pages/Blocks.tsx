import { useQuery } from '@tanstack/react-query';
import { Table, Card } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import { fetchBlocks } from '../services/api';
import { useLocaleStore } from '../stores/useLocaleStore';

interface BlockData {
  number: string;
  hash: string;
  timestamp: string;
  transactions: number;
}

const Blocks = () => {
  const { messages } = useLocaleStore();
  
  const { data, isLoading } = useQuery(['blocks'], fetchBlocks);

  const columns: ColumnsType<BlockData> = [
    {
      title: messages.blocks.columns.number,
      dataIndex: 'number',
      key: 'number',
    },
    {
      title: messages.blocks.columns.hash,
      dataIndex: 'hash',
      key: 'hash',
      ellipsis: true,
    },
    {
      title: messages.blocks.columns.timestamp,
      dataIndex: 'timestamp',
      key: 'timestamp',
    },
    {
      title: messages.blocks.columns.transactions,
      dataIndex: 'transactions',
      key: 'transactions',
    },
  ];

  return (
    <Card title={messages.blocks.title}>
      <Table
        columns={columns}
        dataSource={data}
        loading={isLoading}
        rowKey="number"
        pagination={{ pageSize: 10 }}
      />
    </Card>
  );
};

export default Blocks; 