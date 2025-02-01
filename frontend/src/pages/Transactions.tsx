import { useQuery } from '@tanstack/react-query';
import { Table, Card } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import { fetchTransactions } from '../services/api';
import { useLocaleStore } from '../stores/useLocaleStore';

interface TransactionData {
  hash: string;
  from: string;
  to: string;
  value: string;
  timestamp: string;
}

const Transactions = () => {
  const { messages } = useLocaleStore();
  
  const { data, isLoading } = useQuery(['transactions'], fetchTransactions);

  const columns: ColumnsType<TransactionData> = [
    {
      title: messages.transactions.columns.hash,
      dataIndex: 'hash',
      key: 'hash',
      ellipsis: true,
    },
    {
      title: messages.transactions.columns.from,
      dataIndex: 'from',
      key: 'from',
      ellipsis: true,
    },
    {
      title: messages.transactions.columns.to,
      dataIndex: 'to',
      key: 'to',
      ellipsis: true,
    },
    {
      title: messages.transactions.columns.value,
      dataIndex: 'value',
      key: 'value',
    },
    {
      title: messages.transactions.columns.timestamp,
      dataIndex: 'timestamp',
      key: 'timestamp',
    },
  ];

  return (
    <Card title={messages.transactions.title}>
      <Table
        columns={columns}
        dataSource={data}
        loading={isLoading}
        rowKey="hash"
        pagination={{ pageSize: 10 }}
      />
    </Card>
  );
};

export default Transactions; 