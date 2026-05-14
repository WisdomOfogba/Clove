import { Icon } from '@iconify/react';
import { AdminLayout } from '../../components/layout/AdminLayout';

export function AdminDashboard() {
  return (
    <AdminLayout>
      <div className="space-y-8">
        <div className="flex items-center justify-between bg-white p-6 rounded-[2.5rem] shadow-sm border border-slate-200">
          <div>
              <h1 className="text-3xl font-black tracking-tighter text-dark">Admin Console</h1>
              <p className="text-sm font-medium text-neutral">Global system overview and vitals.</p>
          </div>
          <div className="flex gap-4">
            <button className="bg-slate-50 text-dark px-6 py-3 rounded-2xl text-sm font-bold flex items-center gap-2 hover:bg-slate-100 transition-colors border border-slate-200">
              <Icon icon="solar:download-square-bold-duotone" className="w-5 h-5" /> Export Data
            </button>
            <button className="bg-dark text-white px-6 py-3 rounded-2xl text-sm font-bold flex items-center gap-2 hover:bg-primary hover:text-dark transition-all shadow-xl">
              <Icon icon="solar:settings-bold-duotone" className="w-5 h-5" /> Global Settings
            </button>
          </div>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
          {[
            { label: 'Total Vendors', value: '1,284', grow: '+12%', icon: 'solar:users-group-two-rounded-bold-duotone' },
            { label: 'Pending Review', value: '42', grow: '-5%', icon: 'solar:clock-circle-bold-duotone' },
            { label: 'Verified Today', value: '18', grow: '+3%', icon: 'solar:shield-check-bold-duotone' },
            { label: 'Revenue (24h)', value: '₦4.2M', grow: '+24%', icon: 'solar:wallet-money-bold-duotone' }
          ].map((stat, i) => (
            <div key={stat.label} className="bg-white p-8 rounded-[2.5rem] border border-slate-200 shadow-sm space-y-6">
              <div className="flex items-start justify-between">
                <div className="w-12 h-12 rounded-[1.25rem] bg-slate-50 flex items-center justify-center text-dark relative overflow-hidden">
                  <div className="absolute inset-0 bg-primary/10" />
                  <Icon icon={stat.icon} className="w-6 h-6 relative z-10 text-primary" />
                </div>
                <span className="text-[10px] font-black tracking-widest text-dark bg-slate-100 px-2.5 py-1 rounded-md">{stat.grow}</span>
              </div>
              <div className="space-y-1">
                <p className="text-neutral text-[10px] font-black uppercase tracking-[0.2em]">{stat.label}</p>
                <h3 className="text-3xl font-black text-dark tracking-tight">{stat.value}</h3>
              </div>
            </div>
          ))}
        </div>

        <div className="bg-white rounded-[3rem] border border-slate-200 shadow-sm p-8">
          <div className="flex items-center justify-between mb-8">
            <h2 className="text-xl font-black text-dark">Pending Clove AI Verifications</h2>
            <select className="bg-slate-50 border border-slate-200 text-xs font-bold rounded-xl px-4 py-2 ring-0 outline-none text-dark">
              <option>All Regions</option>
              <option>Lagos</option>
              <option>Abuja</option>
            </select>
          </div>
          <div className="overflow-x-auto">
            <table className="w-full text-left">
              <thead>
                <tr className="text-[10px] uppercase font-bold text-neutral tracking-widest border-b border-slate-100">
                  <th className="pb-4">Vendor</th>
                  <th className="pb-4">Submitted</th>
                  <th className="pb-4">Trust Score</th>
                  <th className="pb-4">Status</th>
                  <th className="pb-4 text-right">Action</th>
                </tr>
              </thead>
              <tbody className="text-sm border-t border-slate-100">
                {[
                  { name: 'Gourmet Greens', date: '2h ago', score: '94%', status: 'Processing' },
                  { name: 'Kebab King', date: '5h ago', score: '62%', status: 'Flagged' },
                  { name: 'Spicy Delight', date: '1d ago', score: '88%', status: 'Ready' },
                ].map((row, i) => (
                  <tr key={i} className="border-b border-slate-50 last:border-0 hover:bg-slate-50/50 transition-colors">
                    <td className="py-4 font-black text-dark text-base">{row.name}</td>
                    <td className="py-4 font-medium text-neutral">{row.date}</td>
                    <td className="py-4">
                      <div className="flex items-center gap-3">
                         <div className="w-16 h-2 bg-slate-100 rounded-full overflow-hidden">
                           <div className="h-full bg-primary" style={{ width: row.score }} />
                         </div>
                         <div>
                            <span className="text-xs font-black text-dark block">{row.score}</span>
                            <span className="text-[8px] font-bold text-primary tracking-widest uppercase">Powered by Clove AI</span>
                         </div>
                      </div>
                    </td>
                    <td className="py-4">
                      <span className={`px-3 py-1 rounded-md text-[10px] font-black uppercase tracking-widest ${row.status === 'Flagged' ? 'bg-red-50 text-red-600' : 'bg-primary/20 text-dark'}`}>{row.status}</span>
                    </td>
                    <td className="py-4 text-right">
                      <button className="text-primary font-bold hover:text-dark px-3 py-1 rounded-lg hover:bg-slate-100 transition-colors">Review</button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </AdminLayout>
  );
}
