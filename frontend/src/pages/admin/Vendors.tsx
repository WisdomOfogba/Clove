import React from 'react';
import { AdminLayout } from '../../components/layout/AdminLayout';
import { Icon } from '@iconify/react';
import { Link } from 'react-router-dom';

export function AdminVendors() {
  const vendors = [
    { id: 1, name: 'Iya Basira Kitchen', location: 'Surulere, Lagos', status: 'Active', score: 99.1, joined: 'Jan 2024' },
    { id: 2, name: 'The Jollof Loft', location: 'Ikeja, Lagos', status: 'Active', score: 98.4, joined: 'Feb 2024' },
    { id: 3, name: 'Spicy Delight', location: 'Abuja', status: 'Suspended', score: 45.0, joined: 'Mar 2024' },
  ];

  return (
    <AdminLayout>
      <div className="space-y-8">
        <div className="flex items-center justify-between bg-white p-6 rounded-[2.5rem] shadow-sm border border-slate-200">
           <div>
               <h1 className="text-3xl font-black tracking-tighter text-dark">Verified Vendors</h1>
               <p className="text-sm font-medium text-neutral">Manage and monitor all active kitchens on CloveDelight.</p>
           </div>
           <div className="flex gap-4">
              <div className="relative">
                 <Icon icon="solar:minimalistic-magnifer-bold-duotone" className="w-5 h-5 absolute left-4 top-1/2 -translate-y-1/2 text-neutral" />
                 <input type="text" placeholder="Search vendors..." className="bg-slate-50 border border-slate-200 text-sm font-bold rounded-2xl pl-12 pr-4 py-3 outline-none focus:ring-2 focus:ring-primary/20 w-64" />
              </div>
           </div>
        </div>

        <div className="bg-white rounded-[3rem] border border-slate-200 shadow-sm overflow-hidden">
           <table className="w-full text-left">
              <thead>
                <tr className="text-[10px] uppercase font-bold text-neutral tracking-widest border-b border-slate-100 bg-slate-50/50">
                  <th className="px-8 py-6">Kitchen Name</th>
                  <th className="px-8 py-6">Location</th>
                  <th className="px-8 py-6">Trust Score</th>
                  <th className="px-8 py-6">Status</th>
                  <th className="px-8 py-6">Joined</th>
                  <th className="px-8 py-6 text-right">Actions</th>
                </tr>
              </thead>
              <tbody className="text-sm border-t border-slate-100">
                 {vendors.map((v) => (
                   <tr key={v.id} className="border-b border-slate-50 last:border-0 hover:bg-slate-50/50 transition-colors">
                      <td className="px-8 py-6 font-black text-dark text-base">{v.name}</td>
                      <td className="px-8 py-6 font-medium text-neutral">{v.location}</td>
                      <td className="px-8 py-6">
                         <div className="flex items-center gap-3">
                           <div className="w-16 h-2 bg-slate-100 rounded-full overflow-hidden">
                             <div className={`h-full ${v.score > 80 ? 'bg-green-500' : 'bg-red-500'}`} style={{ width: `${v.score}%` }} />
                           </div>
                           <span className="text-xs font-black text-dark">{v.score}%</span>
                        </div>
                      </td>
                      <td className="px-8 py-6">
                         <span className={`px-3 py-1.5 rounded-md text-[10px] font-black uppercase tracking-widest ${v.status === 'Active' ? 'bg-green-50 text-green-600' : 'bg-red-50 text-red-600'}`}>
                           {v.status}
                         </span>
                      </td>
                      <td className="px-8 py-6 font-medium text-neutral">{v.joined}</td>
                      <td className="px-8 py-6 text-right">
                         <Link to={`/admin/review-queue`} className="text-primary font-bold hover:text-dark px-4 py-2 rounded-xl hover:bg-slate-100 transition-colors inline-block">View Audit</Link>
                      </td>
                   </tr>
                 ))}
              </tbody>
           </table>
        </div>
      </div>
    </AdminLayout>
  );
}
