// Autogenerated by Thrift Compiler (0.9.2) @generated
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package m3

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type M3 interface { //M3 Metrics Service

	// Emits a batch of metrics.
	//
	// Parameters:
	//  - Batch
	EmitMetricBatch(batch *MetricBatch) (err error)
}

//M3 Metrics Service
type M3Client struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewM3ClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *M3Client {
	return &M3Client{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewM3ClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *M3Client {
	return &M3Client{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Emits a batch of metrics.
//
// Parameters:
//  - Batch
func (p *M3Client) EmitMetricBatch(batch *MetricBatch) (err error) {
	if err = p.sendEmitMetricBatch(batch); err != nil {
		return
	}
	return
}

func (p *M3Client) sendEmitMetricBatch(batch *MetricBatch) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("emitMetricBatch", thrift.ONEWAY, p.SeqId); err != nil {
		return
	}
	args := EmitMetricBatchArgs{
		Batch: batch,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

type M3Processor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      M3
}

func (p *M3Processor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *M3Processor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *M3Processor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewM3Processor(handler M3) *M3Processor {

	self3 := &M3Processor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self3.processorMap["emitMetricBatch"] = &m3ProcessorEmitMetricBatch{handler: handler}
	return self3
}

func (p *M3Processor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x4 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x4.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x4

}

type m3ProcessorEmitMetricBatch struct {
	handler M3
}

func (p *m3ProcessorEmitMetricBatch) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := EmitMetricBatchArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	if err2 = p.handler.EmitMetricBatch(args.Batch); err2 != nil {
		return true, err2
	}
	return true, nil
}

// HELPER FUNCTIONS AND STRUCTURES

type EmitMetricBatchArgs struct {
	Batch *MetricBatch `thrift:"batch,1" json:"batch"`
}

func NewEmitMetricBatchArgs() *EmitMetricBatchArgs {
	return &EmitMetricBatchArgs{}
}

var EmitMetricBatchArgs_Batch_DEFAULT *MetricBatch

func (p *EmitMetricBatchArgs) GetBatch() *MetricBatch {
	if !p.IsSetBatch() {
		return EmitMetricBatchArgs_Batch_DEFAULT
	}
	return p.Batch
}
func (p *EmitMetricBatchArgs) IsSetBatch() bool {
	return p.Batch != nil
}

func (p *EmitMetricBatchArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *EmitMetricBatchArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Batch = &MetricBatch{}
	if err := p.Batch.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Batch, err)
	}
	return nil
}

func (p *EmitMetricBatchArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("emitMetricBatch_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *EmitMetricBatchArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("batch", thrift.STRUCT, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:batch: %s", p, err)
	}
	if err := p.Batch.Write(oprot); err != nil {
		return fmt.Errorf("%T error writing struct: %s", p.Batch, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:batch: %s", p, err)
	}
	return err
}

func (p *EmitMetricBatchArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EmitMetricBatchArgs(%+v)", *p)
}
