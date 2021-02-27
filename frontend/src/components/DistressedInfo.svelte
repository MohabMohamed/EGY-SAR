<div>
    <h2>LKP</h2>
<MaskedInput 
    bind:value={LKP} 
    options={options} 
    name="LKP" 
   
  />
  <h2>DR Position</h2>
  <MaskedInput 
    bind:value={DRPos} 
    options={options} 
    name="DRPos" 
   
  />
  <h2>LKP Distess Date And Time</h2>
  <DatePicker format="%d/%m/%Y" />
  <TimePicker format="%H:%M:%S" />

  <h2>SAR Unit Date And Time</h2>
  <DatePicker format="%d/%m/%Y" />
  <TimePicker format="%H:%M:%S" />

</div>
<script>
import { MaskedInput } from 'svelte-imask';
import { DatePicker,TimePicker } from "attractions";
const options = {
		mask: "(00°00'00'' N), (00°00'00'' E)"
  };
  let LKP ="";
  let DRPos = "";

</script>