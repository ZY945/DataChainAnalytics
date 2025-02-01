import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import zhCN from 'antd/locale/zh_CN';
import enUS from 'antd/locale/en_US';
import zh from '../locales/zh_CN';
import en from '../locales/en_US';

type LocaleType = 'zh_CN' | 'en_US';

interface LocaleState {
  locale: LocaleType;
  antdLocale: typeof zhCN | typeof enUS;
  messages: typeof zh | typeof en;
  setLocale: (locale: LocaleType) => void;
}

export const useLocaleStore = create<LocaleState>()(
  persist(
    (set) => ({
      locale: 'zh_CN',
      antdLocale: zhCN,
      messages: zh,
      setLocale: (locale: LocaleType) =>
        set(() => ({
          locale,
          antdLocale: locale === 'zh_CN' ? zhCN : enUS,
          messages: locale === 'zh_CN' ? zh : en,
        })),
    }),
    {
      name: 'locale-storage',
    }
  )
); 