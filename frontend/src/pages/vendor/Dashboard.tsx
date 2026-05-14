import React from 'react';
import { motion } from 'motion/react';
import { Icon } from '@iconify/react';
import { Link } from 'react-router-dom';
import { Badge } from '../../components/ui/Badge';
import { VendorLayout } from '../../components/layout/VendorLayout';

export function VendorDashboard() {
  return (
    <VendorLayout>
      <div className="space-y-8">
        
        <header className="flex justify-between items-center bg-white p-6 rounded-[2.5rem] border border-slate-200 shadow-sm">
            <div>
              <div className="flex items-center gap-2 mb-1">
                <h2 className="text-2xl font-black text-dark">Welcome back, Iya Basira</h2>
                <div className="w-6 h-6 rounded-full bg-green-100 flex items-center justify-center text-green-600" title="Verified by Clove AI">
                  <Icon icon="solar:shield-check-bold" className="w-4 h-4" />
                </div>
              </div>
              <p className="text-sm font-medium text-slate-500">Your identity and restaurant are fully verified.</p>
              <p className="text-[10px] font-bold text-primary tracking-wide uppercase mt-1">Powered by Clove AI</p>
            </div>
            <div className="flex gap-4">
              <div className="w-12 h-12 rounded-full bg-slate-100 flex items-center justify-center relative">
                  <Icon icon="solar:bell-bold-duotone" className="w-6 h-6 text-dark" />
                  <span className="w-3 h-3 bg-red-500 rounded-full absolute top-2 right-2 border-2 border-white"></span>
              </div>
            </div>
        </header>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div className="bg-dark text-white p-8 rounded-[2.5rem] shadow-[0_20px_50px_rgba(0,0,0,0.15)] relative overflow-hidden">
              <div className="absolute top-0 right-0 w-32 h-32 bg-primary/20 rounded-full blur-3xl" />
              <Icon icon="solar:wallet-money-bold-duotone" className="w-10 h-10 text-primary mb-4" />
              <p className="text-sm font-bold text-white/60 uppercase tracking-widest mb-1">Today's Sales</p>
              <h3 className="text-4xl font-black">₦45,200</h3>
            </div>
            <div className="bg-white p-8 rounded-[2.5rem] border border-slate-200 shadow-sm">
              <Icon icon="solar:box-bold-duotone" className="w-10 h-10 text-neutral mb-4" />
              <p className="text-sm font-bold text-neutral uppercase tracking-widest mb-1">Active Orders</p>
              <h3 className="text-4xl font-black text-dark">12</h3>
            </div>
            <div className="bg-white p-8 rounded-[2.5rem] border border-slate-200 shadow-sm relative overflow-hidden group">
              <Icon icon="solar:shield-check-bold" className="w-10 h-10 text-green-500 mb-4 transition-colors" />
              <p className="text-sm font-bold text-neutral uppercase tracking-widest mb-1">Clove Trust Score</p>
              <div className="flex items-end gap-2">
                <h3 className="text-4xl font-black text-dark">98.5<span className="text-xl text-slate-400">%</span></h3>
              </div>
            </div>
        </div>

        {/* Recent Orders Table */}
        <div className="bg-white rounded-[3rem] p-8 border border-slate-200 shadow-sm">
            <div className="flex justify-between items-center mb-6">
              <h3 className="text-xl font-black text-dark">Incoming Orders</h3>
              <Link to="/vendor/orders" className="text-xs font-bold text-primary uppercase tracking-widest hover:underline">View All</Link>
            </div>
            <div className="overflow-x-auto">
              <table className="w-full text-left">
                  <thead>
                    <tr className="text-[10px] uppercase font-bold text-neutral tracking-widest border-b border-slate-100">
                        <th className="pb-4">Order ID</th>
                        <th className="pb-4">Items</th>
                        <th className="pb-4">Total</th>
                        <th className="pb-4">Status</th>
                        <th className="pb-4 text-right">Action</th>
                    </tr>
                  </thead>
                  <tbody className="text-sm font-bold border-t border-slate-100">
                    {['#8921', '#8922', '#8923'].map((id, i) => (
                      <tr key={i} className="border-b border-slate-50 last:border-0 hover:bg-slate-50/50 transition-colors">
                          <td className="py-4 text-dark">{id}</td>
                          <td className="py-4 text-neutral">{2 + i}x Jollof Rice</td>
                          <td className="py-4 text-dark">₦{(3500 * (2+i)).toLocaleString()}</td>
                          <td className="py-4"><Badge variant="warning">Preparing</Badge></td>
                          <td className="py-4 text-right">
                            <button className="text-primary hover:text-dark px-3 py-1 rounded-lg hover:bg-slate-100 transition-colors">Review</button>
                          </td>
                      </tr>
                    ))}
                  </tbody>
              </table>
            </div>
        </div>

      </div>
    </VendorLayout>
  );
}
