import React from 'react';
import { VendorLayout } from '../../components/layout/VendorLayout';
import { Badge } from '../../components/ui/Badge';

export function VendorOrders() {
  const orders = [
    { id: '#8921', items: '2x Jollof Rice, 1x Coke', total: 7500, status: 'Preparing', date: 'Today, 2:30 PM' },
    { id: '#8920', items: '1x Egusi Soup, 2x Pounded Yam', total: 6000, status: 'Ready', date: 'Today, 1:15 PM' },
    { id: '#8919', items: '3x Asun', total: 15000, status: 'Delivered', date: 'Today, 12:00 PM' },
    { id: '#8918', items: '1x Fried Rice', total: 3500, status: 'Delivered', date: 'Yesterday' },
  ];

  return (
    <VendorLayout>
      <div className="space-y-8">
        <div className="flex justify-between items-center bg-white p-6 rounded-[2.5rem] border border-slate-200 shadow-sm">
            <div>
              <h2 className="text-2xl font-black text-dark">Order History</h2>
              <p className="text-sm font-medium text-neutral">Manage your past and present orders.</p>
            </div>
        </div>

        <div className="bg-white rounded-[3rem] p-8 border border-slate-200 shadow-sm overflow-x-auto">
            <table className="w-full text-left min-w-[600px]">
                <thead>
                  <tr className="text-[10px] uppercase font-bold text-neutral tracking-widest border-b border-slate-100">
                      <th className="pb-4">Order ID / Time</th>
                      <th className="pb-4">Items</th>
                      <th className="pb-4">Total</th>
                      <th className="pb-4">Status</th>
                      <th className="pb-4 text-right">Action</th>
                  </tr>
                </thead>
                <tbody className="text-sm border-t border-slate-100">
                  {orders.map((order) => (
                    <tr key={order.id} className="border-b border-slate-50 last:border-0 hover:bg-slate-50/50 transition-colors">
                        <td className="py-4">
                           <div className="font-bold text-dark">{order.id}</div>
                           <div className="text-xs font-medium text-neutral">{order.date}</div>
                        </td>
                        <td className="py-4 font-medium text-neutral">{order.items}</td>
                        <td className="py-4 font-black text-dark">₦{order.total.toLocaleString()}</td>
                        <td className="py-4">
                           <Badge variant={order.status === 'Delivered' ? 'success' : order.status === 'Ready' ? 'info' : 'warning'}>
                             {order.status}
                           </Badge>
                        </td>
                        <td className="py-4 text-right">
                           {order.status !== 'Delivered' && (
                             <button className="text-primary font-bold hover:text-dark px-3 py-1 rounded-lg hover:bg-slate-100 transition-colors">Update</button>
                           )}
                        </td>
                    </tr>
                  ))}
                </tbody>
            </table>
        </div>
      </div>
    </VendorLayout>
  );
}
