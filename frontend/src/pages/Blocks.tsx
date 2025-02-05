import { useQuery } from '@tanstack/react-query';
import { Table, Card, Tag, Typography } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import { Connection, clusterApiUrl } from '@solana/web3.js';
import { useLocaleStore } from '../stores/useLocaleStore';
import { useState } from 'react';

const { Text } = Typography;

interface BlockData {
  blockhash: string;
  blockHeight: number;
  blockTime: number;
  parentSlot: number;
  previousBlockhash: string;
  transactions: any[];
}

const Blocks = () => {
  const { messages } = useLocaleStore();
  const [connection] = useState(
    new Connection(clusterApiUrl('mainnet-beta'), 'confirmed')
  );

  const { data, isLoading } = useQuery(['blocks'], async () => {
    const slot = await connection.getSlot();
    const blocks: BlockData[] = [];
    
    // 获取最近的10个区块
    for (let i = 0; i < 10; i++) {
      const block = await connection.getBlock(slot - i, {
        maxSupportedTransactionVersion: 0
      });
      if (block) {
        blocks.push({
          blockhash: block.blockhash,
          blockHeight: block.parentSlot + 1,
          blockTime: block.blockTime || 0,
          parentSlot: block.parentSlot,
          previousBlockhash: block.previousBlockhash,
          transactions: block.transactions || [],
        });
      }
    }
    return blocks;
  }, {
    refetchInterval: 10000, // 每10秒刷新一次
  });

  const columns: ColumnsType<BlockData> = [
    {
      title: messages.blocks.columns.number,
      dataIndex: 'blockHeight',
      key: 'blockHeight',
      render: (height: number) => (
        <Text strong>{height.toLocaleString()}</Text>
      ),
    },
    {
      title: messages.blocks.columns.hash,
      dataIndex: 'blockhash',
      key: 'blockhash',
      render: (hash: string) => (
        <Text copyable ellipsis style={{ maxWidth: 200 }}>
          {hash}
        </Text>
      ),
    },
    {
      title: messages.blocks.columns.timestamp,
      dataIndex: 'blockTime',
      key: 'blockTime',
      render: (time: number) => (
        <Text>{new Date(time * 1000).toLocaleString()}</Text>
      ),
    },
    {
      title: messages.blocks.columns.transactions,
      dataIndex: 'transactions',
      key: 'transactions',
      render: (txs: any[]) => (
        <Tag color="blue">{txs.length}</Tag>
      ),
    },
    {
      title: '父区块哈希',
      dataIndex: 'previousBlockhash',
      key: 'previousBlockhash',
      render: (hash: string) => (
        <Text copyable ellipsis style={{ maxWidth: 200 }}>
          {hash}
        </Text>
      ),
    },
  ];

  return (
    <Card 
      title={messages.blocks.title}
      extra={
        <Tag color="green">
          Solana Mainnet
        </Tag>
      }
    >
      <Table
        columns={columns}
        dataSource={data}
        loading={isLoading}
        rowKey="blockhash"
        pagination={false}
      />
    </Card>
  );
};

export default Blocks; 