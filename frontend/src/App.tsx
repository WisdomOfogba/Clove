/**
 * @license
 * SPDX-License-Identifier: Apache-2.0
 */

import { Outlet } from 'react-router-dom';
import { Navbar } from './components/layout/Navbar';
import { Footer } from './components/layout/Footer';
import { motion, AnimatePresence } from 'motion/react';

export default function App() {
  return (
    <div className="min-h-screen flex flex-col font-sans bg-slate-50/30">
      <Navbar />
      <main className="flex-1 flex flex-col pt-4 pb-20">
        <AnimatePresence mode="wait">
          <motion.div
            className="flex-1 flex flex-col"
            key={window.location.pathname}
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0, y: -10 }}
            transition={{ duration: 0.2 }}
          >
            <Outlet />
          </motion.div>
        </AnimatePresence>
      </main>
      <Footer />
    </div>
  );
}
