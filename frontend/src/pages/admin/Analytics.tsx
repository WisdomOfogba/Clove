import React from 'react';
import { AdminLayout } from '../../components/layout/AdminLayout';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, AreaChart, Area } from 'recharts';

export function AdminAnalytics() {
  const data = [
    { name: 'Mon', revenue: 4000, audits: 24 },
    { name: 'Tue', revenue: 3000, audits: 13 },
    { name: 'Wed', revenue: 2000, audits: 98 },
    { name: 'Thu', revenue: 2780, audits: 39 },
    { name: 'Fri', revenue: 1890, audits: 48 },
    { name: 'Sat', revenue: 2390, audits: 38 },
    { name: 'Sun', revenue: 3490, audits: 43 },
  ];

  return (
    <AdminLayout>
      <div className="space-y-8">
        <div className="flex items-center justify-between bg-white p-6 rounded-[2.5rem] shadow-sm border border-slate-200">
           <div>
               <h1 className="text-3xl font-black tracking-tighter text-dark">Data Insights</h1>
               <p className="text-sm font-medium text-neutral">Platform-wide analytics and performance.</p>
           </div>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
           <div className="bg-white p-8 rounded-[3rem] border border-slate-200 shadow-sm space-y-6 flex flex-col h-[400px]">
              <div>
                 <h2 className="text-xl font-black text-dark tracking-tighter">Revenue Overview</h2>
                 <p className="text-xs font-bold text-neutral">Platform fee collections over 7 days.</p>
              </div>
              <div className="flex-1">
                 <ResponsiveContainer width="100%" height="100%">
                    <AreaChart data={data} margin={{ top: 10, right: 10, left: 0, bottom: 0 }}>
                       <defs>
                          <linearGradient id="colorRev" x1="0" y1="0" x2="0" y2="1">
                             <stop offset="5%" stopColor="#4f46e5" stopOpacity={0.3}/>
                             <stop offset="95%" stopColor="#4f46e5" stopOpacity={0}/>
                          </linearGradient>
                       </defs>
                       <CartesianGrid strokeDasharray="3 3" vertical={false} stroke="#f1f5f9" />
                       <XAxis dataKey="name" axisLine={false} tickLine={false} tick={{ fontSize: 10, fontWeight: 700, fill: '#94a3b8' }} />
                       <YAxis axisLine={false} tickLine={false} tick={{ fontSize: 10, fontWeight: 700, fill: '#94a3b8' }} tickFormatter={(val) => `₦${val/1000}k`} />
                       <Tooltip cursor={{ stroke: '#cbd5e1', strokeWidth: 1, strokeDasharray: '4 4' }} contentStyle={{ borderRadius: '1rem', border: 'none', boxShadow: '0 10px 15px -3px rgb(0 0 0 / 0.1)' }} />
                       <Area type="monotone" dataKey="revenue" stroke="#4f46e5" strokeWidth={3} fillOpacity={1} fill="url(#colorRev)" />
                    </AreaChart>
                 </ResponsiveContainer>
              </div>
           </div>

           <div className="bg-white p-8 rounded-[3rem] border border-slate-200 shadow-sm space-y-6 flex flex-col h-[400px]">
              <div>
                 <h2 className="text-xl font-black text-dark tracking-tighter">Automated Audits</h2>
                 <p className="text-xs font-bold text-neutral">AI compliance checks per day.</p>
              </div>
              <div className="flex-1">
                 <ResponsiveContainer width="100%" height="100%">
                    <LineChart data={data} margin={{ top: 10, right: 10, left: -20, bottom: 0 }}>
                       <CartesianGrid strokeDasharray="3 3" vertical={false} stroke="#f1f5f9" />
                       <XAxis dataKey="name" axisLine={false} tickLine={false} tick={{ fontSize: 10, fontWeight: 700, fill: '#94a3b8' }} />
                       <YAxis axisLine={false} tickLine={false} tick={{ fontSize: 10, fontWeight: 700, fill: '#94a3b8' }} />
                       <Tooltip cursor={{ stroke: '#cbd5e1', strokeWidth: 1, strokeDasharray: '4 4' }} contentStyle={{ borderRadius: '1rem', border: 'none', boxShadow: '0 10px 15px -3px rgb(0 0 0 / 0.1)' }} />
                       <Line type="monotone" dataKey="audits" stroke="#10b981" strokeWidth={3} dot={{ r: 4, strokeWidth: 2 }} activeDot={{ r: 6 }} />
                    </LineChart>
                 </ResponsiveContainer>
              </div>
           </div>
        </div>
      </div>
    </AdminLayout>
  );
}
