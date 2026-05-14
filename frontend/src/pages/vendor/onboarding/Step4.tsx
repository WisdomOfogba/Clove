import { useOnboardingStore } from '../../../store';
import { Icon } from '@iconify/react';
import { motion } from 'motion/react';

export function Step4() {
  const { data, setStep, setStatus } = useOnboardingStore();

  const handleFinish = () => {
    setStep(5);
    // Simulate AI Processing delay
    setTimeout(() => {
      setStep(6);
      setStatus('completed');
    }, 4500);
  };

  return (
    <div className="p-8 space-y-8">
      <div className="space-y-1">
        <h2 className="text-xl font-bold text-dark">Review Application</h2>
        <p className="text-sm text-neutral">Double check everything before sending to Clove AI.</p>
      </div>

      <div className="space-y-6">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div className="bg-slate-50 p-6 rounded-2xl border border-slate-100 space-y-4">
            <div className="flex items-center gap-2 text-primary font-bold text-xs uppercase tracking-wider">
              <Icon icon="lucide:shopping-bag" className="w-4 h-4" /> Business
            </div>
            <div className="space-y-2">
              <p className="text-sm font-bold">{data.business.name}</p>
              <p className="text-xs text-neutral">{data.business.address}</p>
              <div className="flex gap-2">
                <span className="text-[10px] bg-white border border-slate-200 px-2 py-0.5 rounded text-neutral">{data.business.category}</span>
                <span className="text-[10px] bg-white border border-slate-200 px-2 py-0.5 rounded text-neutral">{data.business.phone}</span>
              </div>
            </div>
          </div>

          <div className="bg-slate-50 p-6 rounded-2xl border border-slate-100 space-y-4">
            <div className="flex items-center gap-2 text-primary font-bold text-xs uppercase tracking-wider">
              <Icon icon="lucide:shield-check" className="w-4 h-4" /> Legal Proof
            </div>
            <div className="space-y-2">
              <div className="flex justify-between text-xs">
                <span className="text-neutral">CAC Number</span>
                <span className="font-bold">{data.legal.cacNumber}</span>
              </div>
              <div className="flex justify-between text-xs">
                <span className="text-neutral">NIN Number</span>
                <span className="font-bold">{data.legal.ninNumber}</span>
              </div>
              <div className="flex gap-2 pt-2">
                 <div className="w-8 h-8 rounded bg-primary/10 flex items-center justify-center">
                   <Icon icon="lucide:check-circle-2" className="w-4 h-4 text-primary" />
                 </div>
                 <div className="w-8 h-8 rounded bg-primary/10 flex items-center justify-center">
                   <Icon icon="lucide:check-circle-2" className="w-4 h-4 text-primary" />
                 </div>
              </div>
            </div>
          </div>
        </div>

        <div className="bg-primary/5 p-6 rounded-2xl border border-primary/10 space-y-4">
          <div className="flex items-center gap-2 text-primary font-bold text-xs uppercase tracking-wider">
            <Icon icon="lucide:user" className="w-4 h-4" /> Proof of Life
          </div>
          <div className="grid grid-cols-5 gap-2">
            {data.proof.kitchenPhotos.map((img, i) => (
              <img key={i} src={img} className="w-full aspect-square object-cover rounded-lg border border-primary/20" />
            ))}
            {data.proof.mealPhotos.map((img, i) => (
              <img key={i} src={img} className="w-full aspect-square object-cover rounded-lg border border-primary/20" />
            ))}
            <img src={data.proof.selfie!} className="w-full aspect-square object-cover rounded-lg border-2 border-primary ring-2 ring-primary/20" />
          </div>
        </div>
      </div>

      <div className="flex gap-4">
        <button
          type="button"
          onClick={() => setStep(3)}
          className="flex-1 bg-slate-50 text-neutral rounded-xl py-4 font-bold border border-slate-100 hover:bg-slate-100 transition-all flex items-center justify-center gap-2"
        >
          <Icon icon="lucide:arrow-left" className="w-5 h-5" /> Back
        </button>
        <button
          onClick={handleFinish}
          className="flex-[2] bg-primary text-white rounded-xl py-4 font-bold hover:bg-primary-dark transition-all flex items-center justify-center gap-2 shadow-lg shadow-primary/20"
        >
          Submit for AI Verification
          <Icon icon="lucide:shield-check" className="w-5 h-5" />
        </button>
      </div>
    </div>
  );
}
